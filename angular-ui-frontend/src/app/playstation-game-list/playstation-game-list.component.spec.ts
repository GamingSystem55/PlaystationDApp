import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PlaystationGameListComponent } from './playstation-game-list.component';

describe('PlaystationGameListComponent', () => {
  let component: PlaystationGameListComponent;
  let fixture: ComponentFixture<PlaystationGameListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PlaystationGameListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PlaystationGameListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
