import axios from "axios";
import { API_URL } from "./url";

export const auth = async () => {
  try {
    const resp = await axios.get(`${API_URL}/dashboard`);
    console.log(resp.data);
    console.log(resp);
    return resp.data;
  } catch (error) {
    console.log(error);
  }
};
