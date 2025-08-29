import { useState, useEffect } from 'react';
import type { FC } from 'react';
import { Button, Card } from '@telegram-apps/telegram-ui';
import { Page } from '@/components/Page.tsx';
import { hapticFeedback } from '@telegram-apps/sdk';

type CardItem = {
  id: number;
  revealed: boolean;
  prize: string;
  weight: number;
  rarity: 'common' | 'rare' | 'epic' | 'legendary';
  color: string;
};

type GameStats = {
  totalGames: number;
  prizesWon: { [key: string]: number };
};

export const ThemeParamsPage: FC = () => {
  const [cards, setCards] = useState<CardItem[]>([]);
  const [selected, setSelected] = useState<number | null>(null);
  const [gameActive, setGameActive] = useState(false);
  const [animating, setAnimating] = useState(false);
  const [stats, setStats] = useState<GameStats>({ totalGames: 0, prizesWon: {} });
  const [showResult, setShowResult] = useState(false);

  const prizes: CardItem[] = [
    { id: 0, revealed: false, prize: '🎁 Premium Gift', weight: 5, rarity: 'legendary', color: '#FFD700' },
    { id: 1, revealed: false, prize: '💎 x5 Gems', weight: 10, rarity: 'epic', color: '#9C27B0' },
    { id: 2, revealed: false, prize: '⭐ x3 Stars', weight: 15, rarity: 'rare', color: '#2196F3' },
    { id: 3, revealed: false, prize: '🪙 x10 Coins', weight: 25, rarity: 'common', color: '#4CAF50' },
    { id: 4, revealed: false, prize: '🎪 Fun Sticker', weight: 25, rarity: 'common', color: '#FF9800' },
    { id: 5, revealed: false, prize: '🔥 Bonus Round', weight: 20, rarity: 'rare', color: '#F44336' },
  ];

  useEffect(() => {
    const saved = localStorage.getItem('cardGameStats');
    if (saved) {
      setStats(JSON.parse(saved));
    }
  }, []);

  const initCards = () => {
    if (gameActive) return;
    
    const shuffled: CardItem[] = prizes
      .map((c) => ({ ...c, revealed: false }))
      .sort(() => Math.random() - 0.5);

    setCards(shuffled);
    setSelected(null);
    setGameActive(true);
    setShowResult(false);
    setAnimating(true);
    
    hapticFeedback.impactOccurred('medium');
    
    setTimeout(() => setAnimating(false), 600);
  };

  const revealCard = (id: number) => {
    if (!gameActive || selected !== null) return;
    
    setSelected(id);
    hapticFeedback.impactOccurred('heavy');
    
    const selectedCard = cards.find(c => c.id === id);
    if (selectedCard) {
      setStats(prev => {
        const newStats = {
          totalGames: prev.totalGames + 1,
          prizesWon: {
            ...prev.prizesWon,
            [selectedCard.prize]: (prev.prizesWon[selectedCard.prize] || 0) + 1
          }
        };
        localStorage.setItem('cardGameStats', JSON.stringify(newStats));
        return newStats;
      });
    }
    
    setTimeout(() => {
      setCards(prev => prev.map(c => ({ ...c, revealed: true })));
      setShowResult(true);
      setTimeout(() => setGameActive(false), 2000);
    }, 500);
  };

  const getRarityGlow = (rarity: string) => {
    switch (rarity) {
      case 'legendary': return '0 0 20px #FFD700, 0 0 40px #FFD700';
      case 'epic': return '0 0 15px #9C27B0, 0 0 30px #9C27B0';
      case 'rare': return '0 0 10px #2196F3, 0 0 20px #2196F3';
      default: return 'none';
    }
  };

  return (
    <Page>
      <div style={{ padding: '20px', maxWidth: '400px', margin: '0 auto' }}>
        <div style={{ textAlign: 'center', marginBottom: '20px' }}>
          <h2 style={{ margin: '0 0 10px 0', fontSize: '24px', fontWeight: 'bold' }}>🎰 Призовые Карты</h2>
          <div style={{ fontSize: '14px', color: '#666', marginBottom: '15px' }}>
            Игр сыграно: {stats.totalGames}
          </div>
        </div>
        
        <div style={{ textAlign: 'center', marginBottom: '20px' }}>
          <Button 
            onClick={initCards} 
            disabled={gameActive}
            style={{
              background: gameActive ? '#ccc' : 'linear-gradient(45deg, #FF6B6B, #4ECDC4)',
              border: 'none',
              padding: '12px 24px',
              fontSize: '16px',
              fontWeight: 'bold',
              borderRadius: '25px',
              boxShadow: gameActive ? 'none' : '0 4px 15px rgba(0,0,0,0.2)',
              transform: gameActive ? 'none' : 'translateY(-2px)',
              transition: 'all 0.3s ease'
            }}
          >
            {gameActive ? '🎮 Игра идет...' : '🚀 Начать игру'}
          </Button>
        </div>
        
        <div style={{ 
          display: 'grid',
          gridTemplateColumns: 'repeat(3, 1fr)',
          gap: '12px',
          marginTop: '20px',
          perspective: '1000px'
        }}>
          {cards.map((card) => (
            <div
              key={card.id}
              style={{
                width: '100%',
                height: '140px',
                transformStyle: 'preserve-3d',
                transition: 'transform 0.6s ease',
                transform: card.revealed ? 'rotateY(180deg)' : animating ? 'rotateY(10deg) scale(0.95)' : 'none'
              }}
            >
              <Card
                style={{
                  position: 'absolute',
                  width: '100%',
                  height: '100%',
                  backfaceVisibility: 'hidden',
                  cursor: gameActive && selected === null ? 'pointer' : 'default',
                  background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                  border: 'none',
                  borderRadius: '15px',
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                  fontSize: '40px',
                  boxShadow: selected === card.id ? `0 0 25px ${card.color}` : '0 4px 15px rgba(0,0,0,0.1)',
                  transform: selected === card.id ? 'scale(1.05)' : 'scale(1)',
                  transition: 'all 0.3s ease'
                }}
                onClick={() => revealCard(card.id)}
              >
                🎴
              </Card>
              
              <Card
                style={{
                  position: 'absolute',
                  width: '100%',
                  height: '100%',
                  backfaceVisibility: 'hidden',
                  transform: 'rotateY(180deg)',
                  background: `linear-gradient(135deg, ${card.color}22, ${card.color}44)`,
                  border: `2px solid ${card.color}`,
                  borderRadius: '15px',
                  display: 'flex',
                  flexDirection: 'column',
                  alignItems: 'center',
                  justifyContent: 'center',
                  fontSize: '16px',
                  fontWeight: 'bold',
                  color: '#333',
                  textAlign: 'center',
                  padding: '8px',
                  boxShadow: getRarityGlow(card.rarity),
                  animation: showResult && selected === card.id ? 'winner 1s ease-in-out' : 'none'
                }}
              >
                <div style={{ fontSize: '24px', marginBottom: '4px' }}>
                  {card.prize.split(' ')[0]}
                </div>
                <div style={{ fontSize: '12px', lineHeight: '1.2' }}>
                  {card.prize.split(' ').slice(1).join(' ')}
                </div>
                <div style={{
                  fontSize: '10px',
                  marginTop: '4px',
                  padding: '2px 6px',
                  borderRadius: '10px',
                  background: card.color,
                  color: 'white',
                  textTransform: 'uppercase'
                }}>
                  {card.rarity}
                </div>
              </Card>
            </div>
          ))}
        </div>
        
        {showResult && selected !== null && (
          <div style={{
            textAlign: 'center',
            marginTop: '20px',
            padding: '15px',
            background: 'linear-gradient(135deg, #4CAF50, #45a049)',
            borderRadius: '15px',
            color: 'white',
            fontWeight: 'bold',
            animation: 'fadeInUp 0.5s ease-out'
          }}>
            🎉 Поздравляем! Вы выиграли: {cards.find(c => c.id === selected)?.prize}
          </div>
        )}
        
        <style>{`
          @keyframes winner {
            0%, 100% { transform: scale(1) rotateY(180deg); }
            50% { transform: scale(1.1) rotateY(180deg); }
          }
          
          @keyframes fadeInUp {
            from {
              opacity: 0;
              transform: translateY(20px);
            }
            to {
              opacity: 1;
              transform: translateY(0);
            }
          }
        `}</style>
      </div>
    </Page>
  );
};
