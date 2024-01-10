import getToken from "@/app/actions";
import { redirect } from "next/navigation";

export default async function GroupInvitePage({ params }: { params: { groupId: string } }) {
    const token = await getToken();
    if (!token) {
        redirect("/login");
    }

    const res = await fetch(`${process.env.ENDPOINT}/api/groups/invite`, {
        cache: "no-store",
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({group_id: parseInt(params.groupId)})
    })
    const data = await res.json();
    
    return (
        <>
            <h1>{data.error ? data.error : `${process.env.ENDPOINT}/group/${params.groupId}/invite/${data.code}`}</h1>
        </>
    )
}