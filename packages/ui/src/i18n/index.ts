import {
  overwriteGetLocale,
  overwriteSetLocale,
  type Locale,
} from "@/paraglide/runtime.js";

export function initUiI18n(options: {
  getLocale: () => Locale;
  setLocale: (
    locale: Locale,
    options?: { reload?: boolean },
  ) => void | Promise<void>;
}) {
  overwriteGetLocale(options.getLocale);
  overwriteSetLocale(options.setLocale);
}

export function getLocaleDirection(){

}
