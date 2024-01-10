'use server';
import getToken from "@/app/actions";
import { redirect } from "next/navigation";

const createGroup = async (formData: FormData) => {
    const token = await getToken();
    if (!token) {
        redirect("/login");
    }
    
    const request = {
        name: formData.get("name"),
    };

    const response = await fetch(`${process.env.ENDPOINT}/api/groups`, {
        cache: "no-store",
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify(request),
    })
    
    if (response.ok) {
        redirect("/group");
    }

}

export { createGroup };