import {
  getLocale,
  locales,
  setLocale,
  type Locale,
} from "@/paraglide/runtime";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { m } from "#/paraglide/messages";
export function LanguageSwitcher() {
  const currentLocale = getLocale();

  return (
    <Select
      onValueChange={(v) => {
        setLocale(v as Locale);
      }}
    >
      <SelectTrigger className="">
        <SelectValue placeholder={m.locale_lable()} />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {locales.map((locale) => (
            <SelectItem
              key={locale}
              value={locale}
              aria-pressed={locale === currentLocale}
            >
              {m.locale_lable(
                {},
                {
                  locale,
                },
              )}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}
