import React, { createContext, useState } from "react";

export const TodoContext = createContext();

export default function TodoContextProvider({ children }) {
  const [todos, setTodos] = useState(null);

  return (
    <TodoContext.Provider value={{ todos, setTodos }}>
      {children}
    </TodoContext.Provider>
  );
}
