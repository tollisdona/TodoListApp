import { extendTheme, type ThemeConfig } from "@chakra-ui/react";

const config: ThemeConfig ={
    initialColorMode: "dark",
    useSystemColorMode:true,
};

const theme = extendTheme({
    config,
    styles:{
        global: (props:any) => ({
            body: {
                bg: props.colorMode === "dark" ? "" : "gray.400",
                color: props.colorMode === "dark" ? "whiteAlpha.900" : "gray.800",
            },
        }),
    },
});

export default theme;