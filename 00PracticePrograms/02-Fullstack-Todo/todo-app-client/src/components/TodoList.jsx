import {
  Center,
  HStack,
  Spacer,
  StackDivider,
  VStack,
} from "@chakra-ui/layout";
import React, { useContext, useEffect } from "react";
import { Text, Badge } from "@chakra-ui/react";
import { IconButton } from "@chakra-ui/button";
import { useToast } from "@chakra-ui/toast";
import { FaTrash } from "react-icons/fa";
import axios from "axios";
import { BASE_URL } from "../Env";
import { TodoContext } from "../contexts/TodoContextProvider";

export default function TodoList() {
  const { todos, setTodos } = useContext(TodoContext);
  const toast = useToast();

  useEffect(() => {
    getTodos();
  }, []);

  const getTodos = () => {
    axios({
      url: BASE_URL,
      method: "GET",
    })
      .then(({ data }) => {
        setTodos(data);
      })
      .catch((err) => {
        setTodos([]);
        toast({
          title: "Error",
          description: "Unable to fetch todos..",
          status: "error",
          isClosable: true,
        });
      });
  };

  return todos?.length === 0 ? (
    <Center>
      <Badge colorScheme="green" p="4" borderRadius="lg">
        Yay!! No Todos
      </Badge>
    </Center>
  ) : (
    <VStack
      divider={<StackDivider />}
      borderColor="gray.100"
      borderWidth="2px"
      p="4"
      borderRadius="lg"
      width="100%"
      alignItems="stretch"
    >
      {todos?.map((todo, index) => {
        return <TodoListTile key={index} todo={todo} toast={toast} />;
      })}
    </VStack>
  );
}

const TodoListTile = ({ todo, toast }) => {
  const { setTodos } = useContext(TodoContext);

  const removeTodo = (e) => {
    axios({
      url: BASE_URL + todo?.id,
      method: "DELETE",
    })
      .then(({ data }) => {
        setTodos(data);
      })
      .catch((err) => {
        toast({
          title: "Error",
          description: "Unable to delete todo.",
          status: "error",
          isClosable: true,
        });
      });
  };

  return (
    <HStack>
      <Text>{todo.title}</Text>
      <Spacer />
      <IconButton icon={<FaTrash />} onClick={removeTodo} />
    </HStack>
  );
};
