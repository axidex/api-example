import { useState } from 'react';
import { useTonConnectUI, useTonWallet } from '@tonconnect/ui-react';
import { retrieveLaunchParams } from '@telegram-apps/sdk-react';

const createPayloadFromAPI = async (payload: string) => {
    const response = await fetch(`https://axidex.ru:9000/v1/cell?payload=${encodeURIComponent(payload)}`, {
        method: 'POST',
        headers: { 'accept': 'application/json' },
    });

    const result = await response.json();
    if (result.status !== 'SUCCESS') {
        throw new Error('Failed to create cell');
    }
    return result.data;
};

export const useTonPayment = () => {
    const wallet = useTonWallet();
    const [tonConnectUI] = useTonConnectUI();
    const [amount, setAmount] = useState('0.1');

    const handleTopUp = async () => {
        if (!wallet) return;

        const userId = retrieveLaunchParams().tgWebAppData?.user?.id?.toString() ?? '';
        const payloadBOC = await createPayloadFromAPI(userId);

        await tonConnectUI.sendTransaction({
            validUntil: Math.floor(Date.now() / 1000) + 300,
            messages: [{
                address: 'UQAaouYltoqsvhJ1aLDCTQwWYBnRxaCBDr2uzEnbaD_c9O4L',
                amount: BigInt(parseFloat(amount) * 1e9).toString(),
                payload: payloadBOC
            }],
        });
    };

    return { amount, setAmount, handleTopUp };
};