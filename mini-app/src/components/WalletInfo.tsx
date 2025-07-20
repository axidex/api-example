import { FC } from 'react';
import { openLink } from '@telegram-apps/sdk-react';
import { TonConnectButton, Wallet } from '@tonconnect/ui-react';
import { Avatar, Cell, Navigation, Section, Title } from '@telegram-apps/telegram-ui';

interface WalletInfoProps {
    wallet: Wallet;
    className?: string;
}

const isExtendedWallet = (wallet: Wallet): wallet is Wallet & {
    imageUrl: string;
    name: string;
    appName: string;
    aboutUrl: string;
} => {
    return 'imageUrl' in wallet && 'name' in wallet && 'appName' in wallet && 'aboutUrl' in wallet;
};

export const WalletInfo: FC<WalletInfoProps> = ({ wallet, className }) => {
    if (!isExtendedWallet(wallet)) return null;

    return (
        <>
            <Section>
                <Cell
                    before={<Avatar src={wallet.imageUrl} alt="Provider logo" width={60} height={60} />}
                    after={<Navigation>About wallet</Navigation>}
                    subtitle={wallet.appName}
                    onClick={(e) => {
                        e.preventDefault();
                        openLink(wallet.aboutUrl);
                    }}
                >
                    <Title level="3">{wallet.name}</Title>
                </Cell>
            </Section>
            <TonConnectButton className={className} />
        </>
    );
};