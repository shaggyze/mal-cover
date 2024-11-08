# MAL-Cover

Simple API to generate image cover CSS for MyAnimeList list.

## Quick Installation

```bash
docker run -p 34001:34001 shaggyze/mal-cover
```

### Environment Variables

| Name                      |  Default   | Description                                 |
| ------------------------- | :--------: | ------------------------------------------- |
| `MC_APP_PORT`             |  `34001`   | Http server port.                           |
| `MC_APP_READ_TIMEOUT`     |    `1m`    | Http server read timeout.                   |
| `MC_APP_WRITE_TIMEOUT`    |    `1m`    | Http server write timeout.                  |
| `MC_APP_GRACEFUL_TIMEOUT` |   `10s`    | Http server shut down timeout.              |
| `MC_CACHE_DIALECT`        | `inmemory` | Cache type. `nocache`, `redis`, `inmemory`. |
| `MC_CACHE_ADDRESS`        |            | Cache address.                              |
| `MC_CACHE_PASSWORD`       |            | Cache password.                             |
| `MC_CACHE_TIME`           |   `24h`    | Cache time.                                 |

\*_all the environment variables are optional_

## Endpoint

### `/{user}/{type}`

Will generate CSS according to MyAnimeList username and type. For example:

- `http://shaggysmovies.me:34001/shaggyze/anime?style=...`
- `http://shaggysmovies.me:34001/shaggyze/manga?style=...`

### `&{size}=`

Will generate CSS according to MyAnimeList image sizes `small`, `medium`, and `large`. For example:

- `http://shaggysmovies.me:34001/shaggyze/anime?style=...&size=small`
- `http://shaggysmovies.me:34001/shaggyze/manga?style=...&size=medium`

## Styling

**This is the most important part**. The endpoint needs a `style` param. The `style` value depends on how your list show your anime/manga cover image.

For example.

Your list's image cover style is like this.

```css
.animetitle[href*="/37716/"]:before {
  background-image: url(https://cdn.myanimelist.net/images/anime/1889/93555l.jpg);
}
```

Convert it by replacing anime/manga id to `{id}` and image URL to `{url}`.

```css
.animetitle[href*="/{id}/"]:before {
  background-image: url({url});
}
```

Encode it using [URL encode](https://www.urlencoder.org/).

```properties
.animetitle%5Bhref%2A%3D%27%2F%7Bid%7D%2F%27%5D%3Abefore%7Bbackground-image%3Aurl%28%7Burl%7D%29%7D
```

Then use it in endpoint as `style` param.

```
http://shaggysmovies.me:34001/shaggyze/anime?style=.animetitle%5Bhref%2A%3D%27%2F%7Bid%7D%2F%27%5D%3Abefore%7Bbackground-image%3Aurl%28%7Burl%7D%29%7D
```

Good luck.
