import { Link, createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
  component: Home,
})

function Home() {
  return (
    <main style={{ fontFamily: 'system-ui, sans-serif', margin: '2rem auto', maxWidth: 640 }}>
      <h1>Go Backend Template</h1>
      <p>
        TanStack Start SPA served as static assets from the Go binary. API routes stay under{' '}
        <code>/api/v1</code>.
      </p>
      <p>
        <Link to="/login">Sign in</Link>
      </p>
    </main>
  )
}