'use server';

import getToken from "@/app/actions";
import { redirect } from "next/navigation";

const settleUp = async (id: string) => {
    const token = await getToken();
    if (!token) {
        redirect("/login");
    }
    
    const response = await fetch(`${process.env.API_ENDPOINT}/groups/${id}/settlements`, {
        cache: "no-store",
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
    })

    if (!response.ok) {
        return null
    }
    const data = await response.json();
    return data;
}

const comfirmSettlement = async (groupId: string, payerId: string, payeeId: string) => {
    const token = await getToken();

    const response = await fetch(`${process.env.API_ENDPOINT}/settlements/${payerId}/${payeeId}`, {
        cache: "no-store",
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
    })

    if (response.ok) {
        redirect(`/group/${groupId}`);
    }
}

export { settleUp, comfirmSettlement };