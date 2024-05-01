import getToken from "@/app/actions";
import { redirect } from "next/navigation";

export default async function AcceptedGroupInvitePage({ params }: { params: { code: string } }) {
    const token = await getToken();
    if (!token) {
        redirect("/login");
    }

    const res = await fetch(`${process.env.API_ENDPOINT}/invites/${params.code}`, {
        cache: "no-store",
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
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