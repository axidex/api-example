import { FC } from 'react';
import { Button, Input, Section } from '@telegram-apps/telegram-ui';

interface TopUpSectionProps {
    amount: string;
    onAmountChange: (value: string) => void;
    onTopUp: () => void;
}

export const TopUpSection: FC<TopUpSectionProps> = ({ amount, onAmountChange, onTopUp }) => (
    <Section header="TopUp Balance">
        <Input
            type="number"
            min="0.01"
            step="0.01"
            placeholder="Ton"
            value={amount}
            onChange={(e) => onAmountChange(e.target.value)}
        />
        <Button onClick={onTopUp} style={{ marginTop: 12 }}>
            Пополнить
        </Button>
    </Section>
);