'use server';

import getToken from "@/app/actions";

const getGroupName = async (id: string) => {
    const token = await getToken();
    const res = await fetch(`http://golang-dev:8080/groups/${id}`, {
        cache: "no-store",
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        }
    })

    const data = await res.json();
    return data.name;
}

const getGroupMembers = async (id: string) => {
    const token = await getToken();
    const res = await fetch(`http://golang-dev:8080/groups/members/${id}`, {
        cache: "no-store",
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        }
    })

    const data: {id: number, username: string}[] = await res.json();
    return data;
}

const getGroupExpenses = async (id: string) => {
    const token = await getToken();
    const res = await fetch(`http://golang-dev:8080/expenses/${id}`, {
        cache: "no-store",
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        }
    })

    const data = await res.json();
    if (data.error) {
        return []
    }
    return data;
}

export { getGroupName, getGroupMembers, getGroupExpenses };