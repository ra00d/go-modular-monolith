/**
 * Converts snake_case to Pascal_Case (preserving underscores)
 * @example toPascalCase('example_message') => 'Example_Message'
 */
export function toPascalCase(str: string): string {
  return str
    .split("_")
    .filter(Boolean) // Remove empty strings from consecutive underscores
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
    .join("_");
}

/**
 * Converts snake_case to PascalCase with prefix
 * @example toPascalCaseWithPrefix('example_message', 'My') => 'My_Example_Message'
 */
export function toPascalCaseWithPrefix(str: string, prefix: string): string {
  return `${prefix}_${toPascalCase(str)}`;
}
/**
 * Converts snake_case to PascalCase with prefix
 * @example toPascalCaseWithPrefix('example_message', 'My') => 'My_Example_Message'
 */
export function toPascalCaseWithSufix(str: string, sufix: string): string {
  return `${toPascalCase(str)}${sufix}`;
}
/**
 * Converts snake_case to PascalCase (no underscores)
 * @example toPascalCaseNoUnderscore('example_message') => 'ExampleMessage'
 */
export function toPascalCaseNoUnderscore(str: string): string {
  return str
    .split("_")
    .filter(Boolean)
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
    .join("");
}
