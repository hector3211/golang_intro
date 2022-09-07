import { Box, Flex, Text } from "@chakra-ui/react";
import React, { useState, useEffect } from "react";
import axios from "axios";

const SingleBook = () => {
  const [data, setData] = useState([]);
  useEffect(() => {
    const getBookData = async () =>
      await axios
        .get("http://localhost:8000/todo/first")
        .then((response) => {
          setData(response.data.todos);
        })
        .catch((error) => console.log(error));
    getBookData();
  }, []);
  return (
    <Box>
      <Flex direction="column" justify="space-between">
        <Flex justify={"space-around"} direction="row">
          <Text>Title</Text>
          <Text>Description</Text>
          <Text>Completed</Text>
        </Flex>
        {data !== undefined && (
          <Flex direction="row" justify="space-around">
            <Text>{data.Title}</Text>
            <Text>{data.Body}</Text>
            <Text>{data.Completed == true ? "true" : "false"}</Text>
          </Flex>
        )}
      </Flex>
    </Box>
  );
};
export default SingleBook;
