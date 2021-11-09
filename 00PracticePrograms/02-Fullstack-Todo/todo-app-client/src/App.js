import { IconButton } from '@chakra-ui/button';
import { useColorMode } from '@chakra-ui/color-mode';
import { Center, Container, Heading } from '@chakra-ui/layout';
import { FaMoon, FaSun } from 'react-icons/fa';
import './App.css';
import AddTodo from './components/AddTodo';
import TodoList from './components/TodoList';
import TodoContextProvider from './contexts/TodoContextProvider';

function App() {

  const { colorMode, toggleColorMode } = useColorMode()

  return (
    <TodoContextProvider>
      <IconButton icon={colorMode === 'light' ? <FaMoon /> : <FaSun />}
          onClick={e => {
            toggleColorMode()
          }}
          m={4}
          isRound={true} size='lg' alignSelf='flex-end' />
      <Container p={4}>
        <Center><Heading fontWeight='extrabold' size='2xl' bgGradient='linear(to-r, pink.500, pink.300, blue.500)' bgClip='text' mt='4' mb='8' >Todo Application</Heading></Center>
        <br />
        <TodoList />
        <AddTodo />
      </Container>
    </TodoContextProvider>
  );
}

export default App;
