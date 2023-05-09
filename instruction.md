# Step-by-Step Instructions for Using the GitHub Action with Discord

## Step 1
* Create a Discord bot

  * Go to the Discord Developer Portal here
  * Click "New Application"
  * Give it a name and click "Create"
  * On the left sidebar, click "Bot"
  * Click "Add Bot"
  Give your bot a username and click "Create Bot"
  Copy the bot token by clicking "Copy" under "Token"
  Find the Discord channel ID

  Open Discord and go to the channel you want notifications to be sent to
  Enable developer mode in Discord's settings (User Settings -> Appearance -> Developer Mode)
  Right-click the channel and click "Copy ID"

## Step 2
* Save the bot token and channel ID as repository secrets in GitHub
 * Go to your repository on GitHub
* Click "Settings" -> "Secrets"
* Click "New repository secret"
* Name the first secret DISCORD_BOT_TOKEN and paste the bot token you copied earlier into the "Value" field
* Name the second secret DISCORD_CHANNEL_ID and paste the channel ID you copied earlier into the "Value" field
* Click "Add secret"

## Step 3
* Add the GitHub Action to your repository
* In your repository on GitHub, click on the "Actions" tab
* Click "New workflow" and select "set up a workflow yourself"
* Copy the code from the main.yaml file in this repository here
* Paste the code into the text editor on GitHub
* Save the file and name it something like discord-notification.yaml
* Click "Start commit" and then "Commit new file"
Now, whenever the specified event is triggered in your repository, a notification will be sent to the specified Discord channel