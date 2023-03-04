import axios, { AxiosResponse } from "axios";
import { useState, useEffect } from "react";
import { useNavigate } from 'react-router-dom';

const navigate = useNavigate();

export let httpClient = axios.create({});

const useApi = <T>(
  path: string,
  axiosFunc: () => Promise<AxiosResponse<T>>,
  initialState: T,
  handleError: ((res: any) => void) | null = null
): T => {
  const [data, setData] = useState<T>(initialState);
  useEffect(() => {
    const func = async () => {
      const res = await axiosFunc().catch((err) => {
        return err.response;
      });
      if (res.status !== 200) {
        handleError ? handleError(res) : navigate("/error");
      } else {
        setData(res.data);
      }
    };
    func();
  }, []);
  return data;
};

export default useApi;
