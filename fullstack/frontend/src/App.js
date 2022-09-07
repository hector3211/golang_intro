import { Box, Text } from "@chakra-ui/react";
import Todos from "./components/single-book";
function App() {
  return (
    <Box bgColor="gray.800" color={"white"} minH="100vh" maxW="100%">
      <Todos />
    </Box>
  );
}

export default App;
