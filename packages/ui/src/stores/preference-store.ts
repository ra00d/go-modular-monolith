type UserPreferences = {
  theme: "light" | "dark" | "system";
  colorScheme: string;
  language: string;
  direction: "ltr" | "rtl";

  sidebar: {
    collapsed: boolean;
    pinned: boolean;
  };

  density: "compact" | "comfortable" | "spacious";

  table: {
    pageSize: number;
    showFilters: boolean;
    compact: boolean;
  };

  notifications: {
    enabled: boolean;
    sound: boolean;
  };

  dashboard: {
    layout: string;
    showWelcome: boolean;
  };

  accessibility: {
    reducedMotion: boolean;
    highContrast: boolean;
  };
};
