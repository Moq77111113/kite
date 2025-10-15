import { Outlet, getRouteApi } from '@tanstack/solid-router'
import Sidebar from '../components/Sidebar'

const rootRoute = getRouteApi('__root__')

export default function MainLayout() {
  const data = rootRoute.useLoaderData()

  return (
    <div class="flex min-h-screen bg-background">
      <Sidebar templates={data()} />
      <main class="ml-64 flex-1">
        <Outlet />
      </main>
    </div>
  )
}
