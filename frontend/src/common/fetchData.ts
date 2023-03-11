import axios from 'axios';

export class RequestInformation {
  method: string;
  path: string;
  body: any;
  params: Record<string, any>;
  constructor(method: string, path: string, params: any, body: any) {
    this.method = method
    this.path = path
    this.params = params
    this.body = body
  }
}

export const fetchData = async (requestInformation: RequestInformation) => {
  return await axios.get(requestInformation.path).then((response) => {
    return response.data;
  })
}