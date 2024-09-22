'use server';
import { redirect, RedirectType } from "next/navigation";


const handleLogin = async () => {
    const res = await fetch(`${process.env.API_ENDPOINT}/auth/line`, {
      cache: "no-store",
      method: "GET",
    })

    if (res.redirected) {
      redirect(res.url);
    }
  }

export { handleLogin };