'use server';
import { cookies } from "next/headers";

const getToken = async () => {
    const cookie = cookies();
    const token = cookie.get("token");
  
    if (token) {
        return token.value;
    }
    return null
  }

export default getToken;