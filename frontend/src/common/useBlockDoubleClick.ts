import { useState } from 'react';

export const useBlockDoubleClick = <T extends any[]>(
  fn: (...args: T) => void | any | Promise<any>,
): [
  wrapFn: (...args: T) => void | any | Promise<any>,
  processing: boolean,
  unblocking: () => void,
] => {
  const [processing, setProcessing] = useState<boolean>(false);

  const wrapperFunc = (...args: T) => {
    if (processing) {
      // block
      return;
    }
    // start block
    setProcessing(true);

    try {
      const result = fn(...args);
      return result;
    } catch (err) {
      // when catch error. unblock process
      unblocking();
      throw err;
    }
  };

  // When process failed. you can exec this func
  const unblocking = () => {
    setProcessing(false);
  };

  return [wrapperFunc, processing, unblocking];
};
