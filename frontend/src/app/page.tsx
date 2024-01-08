import { Button } from "@mui/material";
import getToken from "./actions";
import { redirect } from "next/navigation";

export default async function Home() {
  
  const token = await getToken();
  if (!token) {
    redirect("/login");
  } else {
    redirect("/group");
  }

  return (
    <div className="flex gap-4 m-10">
    </div>
  )
}
