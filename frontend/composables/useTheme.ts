export default function useTheme() {
    const themeFromStorage = localStorage.getItem("theme");

    const defaultEnabled = themeFromStorage 
        ? themeFromStorage === "dark" 
        : window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches;

    const enabled = useState<boolean | null>('theme', () => defaultEnabled);

    const toggleTheme = () => {
      enabled.value = !enabled.value;
      localStorage.setItem("theme", enabled.value ? "dark" : "light");
      setTheme(enabled.value ? "dark" : "light");
    }

    function setTheme(theme: string | null) {
      if (theme === "dark" || (!theme && window.matchMedia("(prefers-color-scheme: dark)").matches)) {
        document.documentElement.setAttribute("data-theme", "dark");
      } else {
        document.documentElement.removeAttribute("data-theme");
      }
    }

    // Set initial theme based on local storage or default color scheme
    setTheme(themeFromStorage || (defaultEnabled ? "dark" : "light"));

    return {
      enabled,
      toggleTheme,
      setTheme,
    }
}
