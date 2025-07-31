import { Container, Stack } from '@chakra-ui/react'
import Navibar from './components/Navibar'
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'

export const BASE_URL = "http://localhost:5000/api";
function App() {
  return (
    <Stack h="100vh">
      <Navibar/>
      <Container>
        <TodoForm/>
        <TodoList/>
      </Container>      
    </Stack>
  )
}

export default App
