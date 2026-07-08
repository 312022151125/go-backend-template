import { createFileRoute } from '@tanstack/react-router'
import { useState, type FormEvent } from 'react'

export const Route = createFileRoute('/login')({
  component: LoginPage,
})

function LoginPage() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [trust, setTrust] = useState(false)
  const [message, setMessage] = useState<string | null>(null)

  async function onSubmit(e: FormEvent) {
    e.preventDefault()
    setMessage(null)
    const res = await fetch('/api/v1/user/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify({ username, password, trust }),
    })
    const text = await res.text()
    setMessage(`${res.status}: ${text}`)
  }

  return (
    <main style={{ fontFamily: 'system-ui, sans-serif', margin: '2rem auto', maxWidth: 400 }}>
      <h1>Login</h1>
      <form onSubmit={onSubmit}>
        <label style={{ display: 'block', marginBottom: '0.75rem' }}>
          Username
          <input
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            style={{ display: 'block', width: '100%', marginTop: 4 }}
            autoComplete="username"
          />
        </label>
        <label style={{ display: 'block', marginBottom: '0.75rem' }}>
          Password
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            style={{ display: 'block', width: '100%', marginTop: 4 }}
            autoComplete="current-password"
          />
        </label>
        <label style={{ display: 'block', marginBottom: '0.75rem' }}>
          <input
            type="checkbox"
            checked={trust}
            onChange={(e) => setTrust(e.target.checked)}
          />{' '}
          Trust this device
        </label>
        <button type="submit">Sign in</button>
      </form>
      {message ? (
        <pre style={{ marginTop: '1rem', whiteSpace: 'pre-wrap' }}>{message}</pre>
      ) : null}
    </main>
  )
}