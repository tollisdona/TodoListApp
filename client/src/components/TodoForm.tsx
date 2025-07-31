import { BASE_URL } from "../App";
import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useEffect, useRef, useState } from "react";
import { IoMdAdd } from "react-icons/io";

const TodoForm = () =>{
    const [newTodo, setNewTodo] = useState("");
    const queryClient = useQueryClient();
    const {mutate: createTodo,isPending:isCreatng} = useMutation({
        mutationKey:['createTodo'],
        mutationFn: async (e:React.FormEvent) => {
            e.preventDefault();
            if(!newTodo.trim()) return alert("Please enter a task!");
            try {
                const res = await fetch(BASE_URL + `/todos`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ body: newTodo }),
                });
                const data = await res.json();
                if (!res.ok) {
                    throw new Error(data.error || 'Failed to create todo');
                }
                setNewTodo("");
                return data;
            } catch (error) {
                console.error("Error creating todo:", error);
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['todos'] });
        },
        onError: (error:any ) => 
            alert(`Error: ${error instanceof Error ? error.message : 'Something went wrong'}`)
    })
    const inputRef = useRef<HTMLInputElement>(null);
    useEffect(()=>{
        inputRef.current?.focus();
    },[]);
    return (
        <form onSubmit={createTodo}>
            <Flex gap={2}>
               <Input
               type="text"
               value={newTodo}
               onChange={(e) =>setNewTodo(e.target.value)} 
               ref={inputRef}
               />
               <Button
                mx ={2}
                type='submit'
                _active={{
                    transform:"scale(.97)"
                }}>
                    {isCreatng? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
                </Button>
            </Flex>
            </form>
    );
}
export default TodoForm;