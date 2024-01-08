'use client';
import { Button, TextField } from "@mui/material"
import { createGroup } from "./actions"

export default function CreateGroupForm() {
    return (
        <form action={createGroup}>
            <h1>Create group</h1>
            <div>
                <TextField label="name" name="name"/>
            </div>
            <Button type="submit">Create</Button>
        </form>
    )
}