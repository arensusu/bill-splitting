'use client'
import { Button } from "@mui/material";
import getToken from "../actions";
import { redirect } from "next/navigation";
import { handleLogin } from "./actions";

export default async function LoginForm() {

  return (
    <Button variant="contained" onClick={()=>{
      handleLogin()
    }}>Login</Button>
  )
}
