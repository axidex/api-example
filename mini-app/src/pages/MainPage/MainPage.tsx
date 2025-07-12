import {openLink} from '@telegram-apps/sdk-react';
import { TonConnectButton, useTonConnectUI, useTonWallet } from '@tonconnect/ui-react';
import {
    Avatar,
    Button,
    Cell,
    Input,
    List,
    Navigation,
    Placeholder,
    Section,
    Text,
    Title,
} from '@telegram-apps/telegram-ui';
import type { FC } from 'react';
import { useState } from 'react';

import { DisplayData } from '@/components/DisplayData/DisplayData.tsx';
import { Page } from '@/components/Page.tsx';
import { bem } from '@/css/bem.ts';

import "./MainPage.css"

const [, e] = bem('ton-connect-page');

export const MainPage: FC = () => {
    const wallet = useTonWallet();
    const [tonConnectUI] = useTonConnectUI();
    const [amount, setAmount] = useState('0.1'); // default 0.1 TON

    const handleTopUp = async () => {
        if (!wallet) return;

        const response = await tonConnectUI.sendTransaction({
            validUntil: Math.floor(Date.now() / 1000) + 300, // 5 мин
            messages: [
                {
                    address: '0QChdPRtnA0M4a4O1eOMNqp-dO3dxYftquBxyemhDpWAw8DG',
                    amount: BigInt(parseFloat(amount) * 1e9).toString()
                },
            ],
        });

        console.log(response.boc)
    };

    if (!wallet) {
        return (
            <Page>
                <Placeholder
                    className={e('placeholder')}
                    header="TON Connect"
                    description={
                        <>
                            <Text>
                                To display the data related to the TON Connect, it is required to connect your
                                wallet
                            </Text>
                            <TonConnectButton className={e('button')} />
                        </>
                    }
                />
            </Page>
        );
    }

    const {
        account: { chain, publicKey, address },
        device: {
            appName,
            appVersion,
            maxProtocolVersion,
            platform,
            features,
        },
    } = wallet;

    return (
        <Page>
            <List>
                {'imageUrl' in wallet && (
                    <>
                        <Section>
                            <Cell
                                before={
                                    <Avatar src={wallet.imageUrl} alt="Provider logo" width={60} height={60} />
                                }
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
                        <TonConnectButton className={e('button-connected')} />
                    </>
                )}

                <Section header="Пополнение баланса">
                    <Input
                        type="number"
                        min="0.01"
                        step="0.01"
                        placeholder="Сумма в TON"
                        value={amount}
                        onChange={(e) => setAmount(e.target.value)}
                    />
                    <Button onClick={handleTopUp} style={{ marginTop: 12 }}>
                        Пополнить
                    </Button>
                </Section>

                <DisplayData
                    header="Account"
                    rows={[
                        { title: 'Address', value: address },
                        { title: 'Chain', value: chain },
                        { title: 'Public Key', value: publicKey },
                    ]}
                />

                <DisplayData
                    header="Device"
                    rows={[
                        { title: 'App Name', value: appName },
                        { title: 'App Version', value: appVersion },
                        { title: 'Max Protocol Version', value: maxProtocolVersion },
                        { title: 'Platform', value: platform },
                        {
                            title: 'Features',
                            value: features
                                .map(f => typeof f === 'object' ? f.name : undefined)
                                .filter(v => v)
                                .join(', '),
                        },
                    ]}
                />
            </List>
        </Page>
    );
};
