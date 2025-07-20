import { FC } from 'react';
import { useTonWallet } from '@tonconnect/ui-react';
import { List } from '@telegram-apps/telegram-ui';

import { DisplayData } from '@/components/DisplayData/DisplayData.tsx';
import { Page } from '@/components/Page.tsx';
import { bem } from '@/css/bem.ts';
import { useTonPayment } from '@/hooks/useTonPayment';
import { WalletInfo } from '@/components/WalletInfo';
import { TopUpSection } from '@/components/TopUpSection';
import { ConnectWalletPlaceholder } from '@/components/ConnectWalletPlaceholder';

import "./MainPage.css"

const [, e] = bem('ton-connect-page');

export const MainPage: FC = () => {
    const wallet = useTonWallet();
    const { amount, setAmount, handleTopUp } = useTonPayment();

    if (!wallet) {
        return <ConnectWalletPlaceholder className={e('placeholder')} />;
    }

    const { account, device } = wallet;

    return (
        <Page>
            <List>
                <WalletInfo wallet={wallet} className={e('button-connected')} />

                <TopUpSection
                    amount={amount}
                    onAmountChange={setAmount}
                    onTopUp={handleTopUp}
                />

                <DisplayData
                    header="Account"
                    rows={[
                        { title: 'Address', value: account.address },
                        { title: 'Chain', value: account.chain },
                        { title: 'Public Key', value: account.publicKey },
                    ]}
                />

                <DisplayData
                    header="Device"
                    rows={[
                        { title: 'App Name', value: device.appName },
                        { title: 'App Version', value: device.appVersion },
                        { title: 'Max Protocol Version', value: device.maxProtocolVersion },
                        { title: 'Platform', value: device.platform },
                        {
                            title: 'Features',
                            value: device.features
                                .map(f => typeof f === 'object' ? f.name : undefined)
                                .filter(Boolean)
                                .join(', '),
                        },
                    ]}
                />
            </List>
        </Page>
    );
};