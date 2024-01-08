
import getToken from "../actions";
import { redirect } from "next/navigation";
import { Button, List, ListItem, ListItemButton, ListSubheader } from "@mui/material";

const GroupPage = async () => {
    const token = await getToken();
    if (!token) {
        redirect("/login");
    }

    const res = await fetch("http://golang-dev:8080/groups", {
        cache: "no-store",
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        }
    })

    const data = await res.json();
    const groups: any[] = data;

    return (
        <div className="grid grid-col-1 gap-4 m-10">
            <List className="w-48">
                <ListSubheader>Group List</ListSubheader>
                {groups && groups.map((group) => {
                    return (
                        <ListItem key={group.id}>
                            <ListItemButton href={`/group/${group.id}`}>{group.name}</ListItemButton>
                        </ListItem>
                    )
                })}
            </List>
            <Button variant="contained" href="/group/create" className="w-48">
                Create Group
            </Button>
        </div>
    )
}

export default GroupPage;