import '@testing-library/jest-dom'
import { render, screen } from '@testing-library/react'
import Page from '../page'
 
describe('Page', () => {
  it('renders a heading', () => {
    render(<Page />)
 
    const heading = screen.getByTestId('testId')
    expect(heading.innerHTML).toBe('Get started by editing&nbsp;<code class=\"font-mono font-bold\">app/page.tsx</code>')
  })
})