import { Button } from "@mui/material";
import getToken from "../actions";
import { redirect } from "next/navigation";


export default async function LoginForm() {
  const token = await getToken();
  if (token) {
    redirect("/");
  }

  return (
    <Button variant="contained" href={`${process.env.API_ENDPOINT}/auth/line`}>Login</Button>
  )
}
