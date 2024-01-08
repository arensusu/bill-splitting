'use server';
import getToken from "@/app/actions";
import { redirect } from "next/navigation";

const createExpense = async (id: string, formData: FormData) => {
    const token = await getToken();
    if (!token) {
        redirect("/login");
    }
    
    const request = {
        groupId: parseInt(id),
        amount: parseInt(formData.get("amount")?.toString() || "0"),
        date: formData.get("date"),
        description: formData.get("description"),
    };

    const response = await fetch("http://golang-dev:8080/expenses", {
        cache: "no-store",
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify(request),
    })
    
    if (response.ok) {
        redirect(`/group/${id}`);
    }
}

export { createExpense };