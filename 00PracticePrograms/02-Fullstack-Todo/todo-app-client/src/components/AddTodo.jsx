import { Button } from "@chakra-ui/button";
import { Input } from "@chakra-ui/input";
import { Box, HStack } from "@chakra-ui/layout";
import { Spinner } from "@chakra-ui/spinner";
import { useToast } from "@chakra-ui/toast";
import axios from "axios";
import React, { useContext, useState } from "react";
import { TodoContext } from "../contexts/TodoContextProvider";
import { BASE_URL } from "../Env";

export default function AddTodo() {
  const [todoText, setTodoText] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const { setTodos } = useContext(TodoContext);
  const toast = useToast();

  function addTodo(e) {
    e.preventDefault();
    setIsLoading(true);

    if (todoText === "") {
      toast({
        title: "No todo text found",
        description: "Please add todo text.",
        status: "error",
        isClosable: true,
      });
      setIsLoading(false);
      return;
    }

    axios({
      url: `${BASE_URL}`,
      method: "POST",
      data: JSON.stringify({ title: todoText }),
    })
      .then(({ data }) => {
        setTodos(data);
        setIsLoading(false);
      })
      .catch((err) => {
        toast({
          title: "Error!",
          description: err.response?.data,
          status: "error",
          isClosable: true,
        });
        setIsLoading(false);
      });
    setTodoText("");
  }

  return (
    <Box mt={"7"}>
      <form onSubmit={addTodo}>
        <HStack>
          <Input
            onChange={(e) => setTodoText(e.target.value)}
            value={todoText}
            type="text"
            placeholder="Write todo..."
            variant="filled"
          />
          <Button colorScheme={"pink"} disabled={isLoading} type="submit">
            {isLoading ? <Spinner /> : "Add Todo"}
          </Button>
        </HStack>
      </form>
    </Box>
  );
}
