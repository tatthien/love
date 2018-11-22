![https://i.imgur.com/MeVRiOM.gif](https://i.imgur.com/MeVRiOM.gif)

# `/love [emoji]`

A Slack's slash command that shows a heart built from a lot of emojis.

> This project is created for Go Lambda Functions learning purpose. Hosted on Netlify.


## Add slash command to your Slack

- [Create your custom slash command](https://api.slack.com/tutorials/your-first-slash-command) with whatever name you want.
- Add the URL to request when your slash command is run: https://hopeful-bose-b795aa.netlify.com/.netlify/functions/love

## Build it yourself

Just one click to discover the universe  ┐(︶▽︶)┌

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/tatthien/love)

**Manual deploys**

_You neet to install and learn how to use Netlify CLI https://www.netlify.com/docs/cli/_

Clone this repo.

```
git clone https://github.com/tatthien/love.git
```

Deploy the project.

```
cd love/
make build
make deploy
```

If you have a site already on Netlify, you can link the project before `make build`

```
netlify link
```



