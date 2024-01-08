'use client';
import { Button, TextField } from "@mui/material";
import { createExpense } from "./actions";
import { DatePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";

export default function CreateExpenseForm({ params }: { params: { groupId: string } }) {
    const createExpenseWithGroupId = createExpense.bind(null, params.groupId);
    return (
        <form action={createExpenseWithGroupId} className="flex max-w-md flex-col gap-4 m-10">
            <h1>Create expense</h1>
            <div>
                <TextField label="amount" name="amount" />
            </div>
            <div>
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <DatePicker label="date" name="date" format="YYYY-MM-DD" />
                </LocalizationProvider>
            </div>
            <div>
                <TextField label="description" name="description" />
            </div>
            <Button type="submit">Create</Button>
            <Button href={`/group/${params.groupId}`}>Back</Button>
        </form>
        
    )
}