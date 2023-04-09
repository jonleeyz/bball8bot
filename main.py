import logging
import health_check

from bot_token import TOKEN as BOT_TOKEN

def main():
    # initialisation and config setup
    logging.basicConfig(format='%(asctime)s | (%(levelname)s) %(message)s',
        level=logging.DEBUG,
        datefmt='%d/%m/%Y %I:%M:%S %p')

    # health check
    health_check.validateFrameworkVersion()
    health_check.executeHealthCheck(BOT_TOKEN)

if __name__ == '__main__':
    main()
