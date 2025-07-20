import { FC } from 'react';
import { TonConnectButton } from '@tonconnect/ui-react';
import { Placeholder, Text } from '@telegram-apps/telegram-ui';
import { Page } from '@/components/Page.tsx';
import {bem} from "@/css/bem.ts";

interface ConnectWalletPlaceholderProps {
    className?: string;
}

const [, e] = bem('ton-connect-page');

export const ConnectWalletPlaceholder: FC<ConnectWalletPlaceholderProps> = ({ className }) => (
    <Page>
        <Placeholder
            className={className}
            header="TON Connect"
            description={
                <>
                    <Text>
                        To display the data related to the TON Connect, it is required to connect your wallet
                    </Text>
                    <TonConnectButton className={e('button')} />
                </>
            }
        />
    </Page>
);