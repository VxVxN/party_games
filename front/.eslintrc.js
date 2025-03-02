module.exports = {
    env: {
        browser: true,
        es2021: true,
    },
    extends: [
        "eslint:recommended",
        "plugin:react/recommended",
        "plugin:react-hooks/recommended",
        "plugin:@typescript-eslint/recommended",
        "plugin:jsx-a11y/recommended",
        "plugin:import/recommended",
        "plugin:import/typescript",
        "prettier",
    ],
    parser: "@typescript-eslint/parser",
    parserOptions: {
        ecmaFeatures: {
            jsx: true,
        },
        ecmaVersion: "latest",
        sourceType: "module",
    },
    plugins: ["react", "react-hooks", "@typescript-eslint", "jsx-a11y", "import", "prettier"],
    rules: {
        "prettier/prettier": "warn",
        "react/react-in-jsx-scope": "off", // Не нужен в Next.js
        "@typescript-eslint/explicit-module-boundary-types": "off",
        "@typescript-eslint/no-explicit-any": "warn",
        "import/order": ["error", { "newlines-between": "always", "alphabetize": { "order": "asc" } }],
        "jsx-a11y/anchor-is-valid": "off",
    },
    settings: {
        react: {
            version: "detect",
        },
        "import/resolver": {
            typescript: {}, // Поддержка абсолютных импортов TypeScript
        },
    },
};
