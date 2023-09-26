from aiohttp import web


async def handle(request: web.Request):
    return web.Response(text=request.remote)


app = web.Application()
app.add_routes([web.get('/', handle)])


if __name__ == '__main__':
    web.run_app(app)
