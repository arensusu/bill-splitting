import getToken from "@/app/actions";
import { redirect } from "next/navigation";

export default async function AcceptedGroupInvitePage({ params }: { params: { code: string } }) {
    const token = await getToken();
    if (!token) {
        redirect("/login");
    }

    const res = await fetch(`${process.env.ENDPOINT}/groups/members`, {
        cache: "no-store",
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({code: params.code})
    })

    const data = await res.json();
    if (!data.error) {
        redirect(`/group/${data.group_id}`);
    }

    return (
        <>
            <h1>{data.error}</h1>
        </>
    )
}