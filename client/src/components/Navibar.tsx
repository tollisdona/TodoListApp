import { Box, Button, Container, Flex, Text, useColorModeValue, useColorMode } from "@chakra-ui/react";
import {IoMoon} from "react-icons/io5";
import {LuSun} from "react-icons/lu";

export default function Navibar() {
    const {colorMode, toggleColorMode} = useColorMode();

    return(
        <Container maxW={"900px"}>
        <Box bg={useColorModeValue("grey.400", "gray.700")} px={4} my={4} borderRadius={"5"}>
            <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
                {/*Left Side */}
                <Flex 
                    alignItems={"center"}
                    gap={3}
                    display={{base: "none", sm: "flex"}}
                    justifyContent={"flex-start"}
                >
                    <img src='/go.png' alt='logo' width={50} height={50}/>
                    <Text fontSize={"40"}>+</Text>
                    <img src='/gozi.png' alt='logo' width={40} height={40}/>
                    <Text fontSize={"40"}>=</Text>
                    <img src='/vite.svg' alt='logo' width={50} height={50} />
                </Flex>
                {/*Right Side */}
                <Flex
                    alignItems={"center"}
                    gap={3}>
                        <Text fontSize={"lg"} fontWeight={500}>
                            Daily Tasks
                        </Text>
                        <Button  onClick={toggleColorMode}>
                            {colorMode === "light" ? <IoMoon size={20}/> : <LuSun size={20} />}
                        </Button>
                    </Flex>
            </Flex>
        </Box> 
        </Container>
    )
}