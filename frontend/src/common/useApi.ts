import axios, { AxiosResponse } from "axios";
import { useState, useEffect } from "react";
import { useNavigate } from 'react-router-dom';

const navigate = useNavigate();

export let httpClient = axios.create({});

const useApi = async <T>(
  path: string,
  axiosFunc: () => Promise<AxiosResponse<T>>,
  initialState: T,
  handleError: ((res: any) => void) | null = null
): Promise<T> => {
  const [data, setData] = useState<T>(initialState);
  const res = await axiosFunc().catch((err) => {
    return err.response;
  });
  if (res.status !== 200) {
    handleError ? handleError(res) : navigate("/error");
  } else {
    setData(res.data);
  }
  return data;
};

export default useApi;
