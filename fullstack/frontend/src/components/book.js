import { Box, Button, Text } from "@chakra-ui/react";
import React from "react";

const Book = (bookData, setBookPrice, deleteBook, setBook) => {
  return (
    <Box>
      <Flex direction="column">
        {bookData !== undefined && (
          <Flex>
            <Text>Title:{bookData.title}</Text>
            <Text>Description:{bookData.description}</Text>
            <Text>Price ${bookData.price}</Text>
            <Button onClick={() => setBookPrice(bookData.id)}>
              Update Price
            </Button>
            <Button onClick={() => setBook(bookData.id)}>
              Update book feilds
            </Button>
            <Button onClick={() => deleteBook(bookData.id)}>Delete</Button>
          </Flex>
        )}
      </Flex>
    </Box>
  );
};

export default Book;
