import getToken from "@/app/actions";
import { redirect } from "next/navigation";

export default async function GroupInvitePage({ params }: { params: { groupId: string } }) {
    const token = await getToken();
    if (!token) {
        redirect("/login");
    }

    const res = await fetch(`${process.env.API_ENDPOINT}/groups/${params.groupId}/invites`, {
        cache: "no-store",
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
    })
    const data = await res.json();
    
    return (
        <>
            <h1>{data.error ? data.error : `${process.env.API_ENDPOINT}/group/${params.groupId}/invite/${data.code}`}</h1>
        </>
    )
}