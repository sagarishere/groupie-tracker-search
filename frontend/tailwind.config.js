/** @type {import('tailwindcss').Config} */
export const content = [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
];
export const theme = {
    extend: {
        typography: {},
    },
};
export const darkMode = ['class', '[data-theme="dark"]'];
export const plugins = [
    require('@tailwindcss/typography'),
];
