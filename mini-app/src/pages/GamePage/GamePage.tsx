import { useState, useEffect } from 'react';
import type { FC } from 'react';
import { Button, Card } from '@telegram-apps/telegram-ui';
import { Page } from '@/components/Page.tsx';
import { hapticFeedback } from '@telegram-apps/sdk';
import './GamePage.css';

// https://cdn.changes.tg/gifts/models/
type CardItem = {
  id: number;
  revealed: boolean;
  prize: Prize;
};

type Prize = {
  id: string;
  name: string;
  rarity: 'common' | 'rare' | 'epic' | 'legendary';
  color: string;
  emoji: string;
};

type GameSession = {
  sessionId: string;
  cards: CardItem[];
  selectedCardId: number;
};

type GameStats = {
  totalGames: number;
  prizesWon: { [key: string]: number };
};

const mockCreateGameSession = async (): Promise<{ sessionId: string; cards: { id: number; revealed: boolean }[] }> => {
  await new Promise(resolve => setTimeout(resolve, 300));
  return {
    sessionId: Math.random().toString(36).substr(2, 9),
    cards: Array.from({ length: 6 }, (_, i) => ({ id: i, revealed: false }))
  };
};

const mockRevealCard = async (sessionId: string, cardId: number): Promise<GameSession> => {
  await new Promise(resolve => setTimeout(resolve, 500));
  
  const allPrizes: Prize[] = [
    { id: 'gift', name: 'Premium Gift', rarity: 'legendary', color: '#FFD700', emoji: '🎁' },
    { id: 'gems', name: 'x5 Gems', rarity: 'epic', color: '#9C27B0', emoji: '💎' },
    { id: 'stars', name: 'x3 Stars', rarity: 'rare', color: '#2196F3', emoji: '⭐' },
    { id: 'coins', name: 'x10 Coins', rarity: 'common', color: '#4CAF50', emoji: '🪙' },
    { id: 'sticker', name: 'Fun Sticker', rarity: 'common', color: '#FF9800', emoji: '🎪' },
    { id: 'bonus', name: 'Bonus Round', rarity: 'rare', color: '#F44336', emoji: '🔥' },
  ];
  
  // Генерируем случайные призы для всех карт
  const cardPrizes = Array.from({ length: 6 }, () => 
    allPrizes[Math.floor(Math.random() * allPrizes.length)]
  );
  
  
  return {
    sessionId,
    cards: cardPrizes.map((prize, i) => ({ 
      id: i, 
      revealed: true, 
      prize 
    })),
    selectedCardId: cardId
  };
};

export const GamePage: FC = () => {
  const [initialSession, setInitialSession] = useState<{ sessionId: string; cards: { id: number; revealed: boolean }[] } | null>(null);
  const [gameSession, setGameSession] = useState<GameSession | null>(null);
  const [selected, setSelected] = useState<number | null>(null);
  const [gameActive, setGameActive] = useState(false);
  const [animating, setAnimating] = useState(false);
  const [stats, setStats] = useState<GameStats>({ totalGames: 0, prizesWon: {} });
  const [showResult, setShowResult] = useState(false);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const saved = localStorage.getItem('cardGameStats');
    if (saved) {
      setStats(JSON.parse(saved));
    }
  }, []);

  const initCards = async () => {
    if (gameActive || loading) return;
    
    setLoading(true);
    setSelected(null);
    setShowResult(false);
    setAnimating(true);
    
    hapticFeedback.impactOccurred('medium');
    
    try {
      const session = await mockCreateGameSession();
      setInitialSession(session);
      setGameSession(null);
      setGameActive(true);
      setTimeout(() => setAnimating(false), 600);
    } catch (error) {
      console.error('Failed to create game session:', error);
    } finally {
      setLoading(false);
    }
  };

  const revealCard = async (id: number) => {
    if (!gameActive || selected !== null || !initialSession || loading) return;
    
    setSelected(id);
    setLoading(true);
    hapticFeedback.impactOccurred('heavy');
    
    try {
      const result = await mockRevealCard(initialSession.sessionId, id);
      setGameSession(result);
      
      const winningPrize = result.cards[result.selectedCardId].prize;
      setStats(prev => {
        const prizeName = `${winningPrize.emoji} ${winningPrize.name}`;
        const newStats = {
          totalGames: prev.totalGames + 1,
          prizesWon: {
            ...prev.prizesWon,
            [prizeName]: (prev.prizesWon[prizeName] || 0) + 1
          }
        };
        localStorage.setItem('cardGameStats', JSON.stringify(newStats));
        return newStats;
      });
      
      setTimeout(() => {
        setShowResult(true);
        setTimeout(() => setGameActive(false), 2000);
      }, 100);
    } catch (error) {
      console.error('Failed to reveal card:', error);
    } finally {
      setLoading(false);
    }
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
      <div className="game-container">
        <div className="game-header">
          <h2 className="game-title">🎰 Призовые Карты</h2>
          <div className="game-stats">
            Игр сыграно: {stats.totalGames}
          </div>
        </div>
        
        <div className="game-controls">
          <Button 
            onClick={initCards} 
            disabled={gameActive || loading}
            className="game-button"
          >
            {loading ? '⏳ Загрузка...' : gameActive ? '🎮 Игра идет...' : '🚀 Начать игру'}
          </Button>
        </div>
        
        <div className="cards-grid">
          {(gameSession?.cards || initialSession?.cards)?.map((card) => (
            <div
              key={card.id}
              className={`card-container ${
                card.revealed ? 'revealed' : ''
              } ${
                animating ? 'animating' : ''
              }`}
            >
              <Card
                className={`card-front ${
                  gameActive && selected === null && !loading ? 'clickable' : ''
                } ${
                  selected === card.id ? 'selected' : ''
                }`}
                onClick={() => revealCard(card.id)}
              >
                🎴
              </Card>
              
              {gameSession && (card as CardItem).prize && (
                <Card
                  className={`card-back ${
                    showResult && card.id === gameSession.selectedCardId ? 'winning' : 'non-winning'
                  }`}
                  style={{
                    background: `linear-gradient(135deg, ${(card as CardItem).prize.color}22, ${(card as CardItem).prize.color}44)`,
                    border: `2px solid ${(card as CardItem).prize.color}`,
                    boxShadow: card.id === gameSession.selectedCardId ? getRarityGlow((card as CardItem).prize.rarity) : 'none'
                  }}
                >
                  <div className="prize-emoji">
                    {(card as CardItem).prize.emoji}
                  </div>
                  <div className="prize-name">
                    {(card as CardItem).prize.name}
                  </div>
                  <div 
                    className="prize-rarity"
                    style={{ background: (card as CardItem).prize.color }}
                  >
                    {(card as CardItem).prize.rarity}
                  </div>
                </Card>
              )}
            </div>
          ))}
        </div>
        
        {showResult && gameSession && (
          <div className="result-banner">
            🎉 Поздравляем! Вы выиграли: {gameSession.cards[gameSession.selectedCardId].prize.emoji} {gameSession.cards[gameSession.selectedCardId].prize.name}
          </div>
        )}
        
      </div>
    </Page>
  );
};
