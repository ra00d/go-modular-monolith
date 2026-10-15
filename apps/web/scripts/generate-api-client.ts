import fs from "node:fs";
import openapiTS, { astToString } from "openapi-typescript";

const ast = await openapiTS(
  new URL("http://localhost:8000/api-docs-json", import.meta.url),
  {
    dedupeEnums: true,
    rootTypes: true,
    rootTypesNoSchemaPrefix: true,
    excludeDeprecated: true,
    defaultNonNullable: true,
    propertiesRequiredByDefault: true,
    exportType: true,
    pathParamsAsTypes: true,
  },
);

const contents = astToString(ast);

// (optional) write to file
fs.writeFileSync("src/types/api-schema.ts", contents);
