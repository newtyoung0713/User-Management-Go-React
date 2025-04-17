# User Management System - React Frontend

A modern React frontend for the User Management System, built with TypeScript, Vite, and Shadcn UI.

## Features

- 🔐 JWT Authentication
- 👥 User Management
- 🎨 Modern UI with Shadcn UI
- 📱 Responsive Design
- 🌐 TypeScript Support
- ⚡ Vite for Fast Development

## Prerequisites

- Node.js >= 16
- pnpm >= 8

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   pnpm install
   ```
3. Copy environment file:
   ```bash
   cp .env.example .env
   ```
4. Start development server:
   ```bash
   pnpm run dev
   ```

## Available Scripts

- `pnpm run dev` - Start development server
- `pnpm run build` - Build for production
- `pnpm run preview` - Preview production build
- `pnpm run lint` - Run ESLint
- `pnpm run test` - Run tests

## Project Structure

```
src/
├── api/          # API client and endpoints
├── components/   # Reusable components
├── contexts/     # React contexts
├── hooks/        # Custom hooks
├── layouts/      # Layout components
├── pages/        # Page components
├── services/     # Business logic
├── types/        # TypeScript types
└── utils/        # Utility functions
```

## Environment Variables

- `VITE_API_BASE_URL` - Backend API base URL
- `VITE_JWT_STORAGE_KEY` - Local storage key for JWT token
- `VITE_ENABLE_MOCK_API` - Enable mock API for development

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License.
