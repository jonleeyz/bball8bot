import logging
import health_check

def main():
    # initialisation and config setup
    logging.basicConfig(level=logging.INFO)
    from bot_token import TOKEN as BOT_TOKEN

    # health check
    health_check.checkFrameworkVersion()
    health_check.executeHealthCheck(BOT_TOKEN)

if __name__ == '__main__':
    main()
