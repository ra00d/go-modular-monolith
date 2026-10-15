import { resolve } from "node:path";
import { defineConfig } from "vite";
import reactPlugin from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";
import dts from "unplugin-dts/vite";

const config = defineConfig({
  resolve: { tsconfigPaths: true },
  plugins: [
    reactPlugin(),
    dts({
      tsconfigPath: "./tsconfig.json",
      insertTypesEntry: true,
      //entryRoot: "src",
    }),
    tailwindcss(),
  ],
  build: {
    lib: {
      entry: {
        components: resolve(import.meta.dirname, "src/components/index.ts"),
        hooks: resolve(import.meta.dirname, "src/hooks/index.ts"),
        lib: resolve(import.meta.dirname, "src/lib/index.ts"),
        types: resolve(import.meta.dirname, "src/types/index.ts"),
        pages: resolve(import.meta.dirname, "src/pages/index.ts"),

        // styles: resolve(import.meta.dirname, "src/styles.css"),
      },
      name: "UI",
      formats: ["es"],
    },
    rolldownOptions: {
      // make sure to externalize deps that shouldn't be bundled
      // into your library
      external: [
        "react",
        "react-dom",
        /^react\//,
        /^react-dom\//,
        /@tanstack*/,
        /@base-ui*/,
      ],
      // output: {
      //   // Provide global variables to use in the UMD build
      //   // for externalized deps
      //   globals: {
      //     react: "React",
      //   },
      // },
    },
  },
});

export default config;
